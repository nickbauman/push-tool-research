(ns async-httpkit.client4
  (:require [clojure.core.async :refer [chan <! <!! >! close! go go-loop]]
            [org.httpkit.client :as http]))

(defonce REQUESTS 1000000)

(defn pmap
  "Like map, except f is applied in parallel. Semi-lazy in that the
  parallel computation stays ahead of the consumption, but doesn't
  realize the entire result unless required. Only useful for
  computationally intensive functions where the time of f dominates
  the coordination overhead."
  {:added "1.0"
   :static true}
  ([f coll]
   (let [n 1024
         rets (map #(future (f %)) coll)
         step (fn step [[x & xs :as vs] fs]
                (lazy-seq
                 (if-let [s (seq fs)]
                   (cons (deref x) (step xs (rest s)))
                   (map deref vs))))]
     (step rets (drop n rets))))
  ([f coll & colls]
   (let [step (fn step [cs]
                (lazy-seq
                 (let [ss (map seq cs)]
                   (when (every? identity ss)
                     (cons (map first ss) (step (map rest ss)))))))]
     (pmap #(apply f %) (step (cons coll colls))))))


(defn process-results
  [result-map response]
  (let [status (:status response)
        curr-count (get result-map status 0)]
    (assoc result-map status (inc curr-count))))

(defn bench
  [url]
  (let [start-time (System/currentTimeMillis)
        responses (pmap deref (repeatedly REQUESTS #(http/get url)))]
    (let [results (reduce process-results {} responses)
          duration (double (/ (- (System/currentTimeMillis) start-time) 1000))]
      (println "Done!")
      (println "Results:" results)
      (println "Req/Sec:" (/ (reduce + (vals results)) duration))
      (println "That took:" duration "seconds"))))

(defn bench-n
  [num-runs url]
  (println "client2 starting" num-runs "runs")
  (dotimes [_ num-runs]
    (bench url)))
