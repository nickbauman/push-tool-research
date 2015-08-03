(ns async-httpkit.client2
  (:require [clojure.core.async :refer [chan <!! >! close! go]]
            [org.httpkit.client :as http]))

(defonce REQUESTS 1000000)
(defonce WORKERS 1024)

(defn run-requests
  [work-chan url]
  (go (dotimes [_ REQUESTS]
        (>! work-chan (http/get url)))
      (close! work-chan)))

(defn bench
  [url]
  (let [start-time (System/currentTimeMillis)
        work-chan (chan WORKERS)]
    (run-requests work-chan url)
    (loop [response (<!! work-chan)
           results {}]
      (if response
        (let [status (:status @response)
              old-result (get results status 0)]
          (recur (<!! work-chan) (assoc results status (inc old-result))))
        (let [duration (double (/ (- (System/currentTimeMillis) start-time) 1000))]
          (println "Done!")
          (println "Results:" results)
          (println "Req/Sec:" (/ (reduce + (vals results)) duration))
          (println "That took:" duration "seconds"))))))

(defn bench-n
  [num-runs url]
  (println "client2 starting" num-runs "runs")
  (dotimes [_ num-runs]
    (bench url)))
