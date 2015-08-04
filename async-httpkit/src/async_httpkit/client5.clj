(ns async-httpkit.client5
  (:require [clojure.core.async :refer [chan <! <!! >! close! go go-loop]]
            [clj-http.client :as http]))

(defonce REQUESTS 100000)
(defonce WORKERS 256)
(defonce WORKSIZE 1024)

(defonce POOL (clj-http.conn-mgr/make-reusable-conn-manager {:timeout 120 :threads 256 :default-per-route 256}))

(defn run-requests
  [work-chan url]
  (go (dotimes [_ REQUESTS]
        (>! work-chan url))
      (close! work-chan)))

(defn run-worker
  [work-chan result-chan]
  (go-loop [url (<! work-chan)]
    (when-not (nil? url)
      (>! result-chan
          (http/get url {:connection-manager POOL}))
      (recur (<! work-chan)))))

(defn process-results
  [result-chan]
  (go-loop [result-response (<! result-chan)
            total-results 1
            result-map {}]
    (let [status (:status result-response)
          old-result (get result-map status 0)
          new-result (assoc result-map status (inc old-result))]
      (if (< total-results REQUESTS)
        (recur (<! result-chan)
               (inc total-results)
               new-result)
        new-result))))

(defn bench
  [url]
  (let [start-time (System/currentTimeMillis)
        work-chan (chan WORKSIZE)
        result-chan (chan WORKSIZE)]
    (run-requests work-chan url)
    (dotimes [_ WORKERS]
      (run-worker work-chan result-chan))
    (let [results (<!! (process-results result-chan))
          duration (double (/ (- (System/currentTimeMillis) start-time) 1000))]
      (println "Done!")
      (println "Results:" results)
      (println "Req/Sec:" (/ (reduce + (vals results)) duration))
      (println "That took:" duration "seconds"))))

(defn bench-n
  [num-runs url]
  (println "client2 starting" num-runs "runs")
  (dotimes [_ num-runs]
    (bench url))
  (clj-http.conn-mgr/shutdown-manager POOL))
