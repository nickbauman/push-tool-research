# async-tornado

A web client and web server that are benched for testing.

## Usage

Run tornado server for testing on port 8080

    $ python torasynch_server.py

Run benmark on tornado AsyncHTTTPClient client connecting to server at URL:

    $ python torasynch_client_bench.py

Run benmark on grequests client connecting to server at URL:

    $ python grequests_client_bench.py

## License

Copyright © 2015 Dang Vang, Nick Bauman

Distributed under the Eclipse Public License either version 1.0 or (at
your option) any later version.
