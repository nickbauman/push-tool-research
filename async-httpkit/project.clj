(defproject async-httpkit "0.1.0-SNAPSHOT"
  :description "Benchmarks for Clojure Research"
  :url "http://example.com/FIXME"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.7.0"]
                 [http-kit "2.1.18"]
                 [org.clojure/core.async "0.1.346.0-17112a-alpha"]
                 [clj-http "2.0.0"]]
  :main async-httpkit.core
  :jvm-opts ["-Xmx1G" "-Xms1G"])
