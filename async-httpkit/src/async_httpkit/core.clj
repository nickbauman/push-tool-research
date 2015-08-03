(ns async-httpkit.core
  (:require [async-httpkit.client :as client]
            [async-httpkit.client2 :as client2]
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
      "server" (server/start)
      "unknown option")))
