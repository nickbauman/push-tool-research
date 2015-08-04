(ns async-httpkit.core
  (:require [async-httpkit.client :as client]
            [async-httpkit.client2 :as client2]
            [async-httpkit.client3 :as client3]
            [async-httpkit.client4 :as client4]
            [async-httpkit.client5 :as client5]
            [async-httpkit.server :as server]))

(defn -main
  [& args]
  (if (seq args)
    (case (first args)
      "client" (if (second args)
                 (client/bench-n 3 (second args))
                 (println "client can't connect to nil endpoint"))
      "client2" (if-let[url (second args)]
                  (client2/bench-n 3 url)
                  (println "client can't connect to nil endpoint"))
      "client3" (if-let[url (second args)]
                  (client3/bench-n 3 url)
                  (println "client can't connect to nil endpoint"))
      "client4" (if-let[url (second args)]
                  (client4/bench-n 3 url)
                  (println "client can't connect to nil endpoint"))
      "client5" (if-let[url (second args)]
                  (client5/bench-n 3 url)
                  (println "client can't connect to nil endpoint"))
      "server" (server/start)
      "unknown option")))
