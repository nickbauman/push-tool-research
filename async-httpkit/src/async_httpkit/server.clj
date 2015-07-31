(ns async-httpkit.server
  (:require [clojure.core.async :refer [chan <!! <! >! close! go go-loop] :as async]
            [org.httpkit.server :as server]))

(defn app [req]
  ;(print ".")
  {:status  200
   :headers {"Content-Type" "text/html"}
   :body    "hello world!"})

(defn start
  []
  (println "starting server")
  (server/run-server app {:port 8080}))
