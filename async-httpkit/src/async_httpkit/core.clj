(ns async-httpkit.core
  (:require [async-httpkit.client :as client]
            [async-httpkit.server :as server]))

(defn -main
  [& args]
  (if (seq args)
    (case (first args)
      "client" (if (second args)
                 (client/benchN 3 (second args))
                 (println "client can't connect to nil endpoint"))
      "server" (server/start)
      "unknown option")))
