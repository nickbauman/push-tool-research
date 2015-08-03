(ns async-httpkit.server
  (:require [org.httpkit.server :as server]))

(defn app [req]
  {:status  200
   :headers {"Content-Type" "text/html"}
   :body    "hello world!"})

(defn start
  []
  (println "starting server")
  (server/run-server app {:port 8080}))
