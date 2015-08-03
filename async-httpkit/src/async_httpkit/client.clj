(ns async-httpkit.client
  (:require [clojure.core.async :refer [chan <!! <! >! close! go go-loop] :as async]
            [org.httpkit.client :as http]))

; for clojure http-kit client benchmarks

(defonce REQUESTS 1000000)
(defonce CONNECTIONS 1024)
(defonce WORKERS 4)

(defn run-requests
  [work-chan, url]
  (go (dotimes [_ REQUESTS]
        (>! work-chan (http/get url)))
      (close! work-chan)))

(defn process-responses
  [work-chan]
  (go-loop [response (<! work-chan)
            results {}]
    (if response
      (let [status (:status @response)
            old-result (get results status 0)]
        (recur (<! work-chan) (assoc results status (inc old-result))))
      results)))

(defn bench
  [url]
  (let [start-time (System/currentTimeMillis)
        work-chan (chan CONNECTIONS)
        _ (run-requests work-chan url)
        all-workers (doall (repeatedly WORKERS #(process-responses work-chan)))
        all-results (<!! (async/into [] (async/merge all-workers)))
        duration (double (/ (- (System/currentTimeMillis) start-time) 1000))
        results (apply merge-with + all-results)]
    (println "Done!")
    (println "Results:" results)
    (println "Req/Sec:" (/ (reduce + (vals results)) duration))
    (println "That took:" duration "seconds")))

(defn bench-n
  [num-runs url]
  (println "client starting" num-runs "runs")
  (dotimes [_ num-runs]
    (bench url)))
