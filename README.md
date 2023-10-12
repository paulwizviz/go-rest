
# Overview

A collection of resources and working examples illustrating concepts and practices designing and implementing RESTful API.

## Concepts

Application Programming Interface (API) is a means for two computational entities to communicate with each other. REST, which stands for REpresentational State Transfer, is a **pattern** of communication over a networking protocol known as HTTP. NOTE: the term **pattern**. REST is **not** a protocol more like a convention for organising communications around Uniform Resource Identifier (URI).

The URIs identify resources as a basis for communication between two computers. In other words, one computer -- i.e. the client -- communicate the other -- i.e. the server -- on resources resource. REST group resources by nouns not verbs, for example:

* like this - `http://localhost:8080/products`;
* but not like this - `http://localhost:8080/getproducts`.

The second example is a Remote Procedure Call (RPC) style API communication pattern.

The interactions between clients and servers are typically based on HTTP Request and Response basis. A request is preceeded by the these HTTP verbs `POST`, `GET`, `PUT` and `DELETE` to server to perform the following operations `CREATE`,  `READ`, `UPDATE` and `DELETE` specificied resources respectively. Here is an example of the a call with verbs:

```
   [client] ----  POST /products HTTP ---> [server] ====> [products data store] 
```

A HTTP request may also include a body in JSON format representing request expected the resource.

The HTTP request returns at least three levels of standard response codes:

* 200-level (Success),
* 400-level (Something is wrong with the request), or
* 500-level (Something is wrong with the server).

A response will also include data in JSON format representing data about the resource and other meta data such as suggesting to client the related on URIs (e.g. HATEOAS).


## Useful resources

* [Architectural Styles and the Design of Network-based Software Architectures - Roy Fielding Dissertation](https://ics.uci.edu/~fielding/pubs/dissertation/top.htm)
* [What Is REST API? Examples And How To Use It: Crash Course System Design #3](https://www.youtube.com/watch?v=-mN3VyJuCjM)

## Disclaimer

* The content of this project are intended for educational purpose only.
* The content is constantly updated and items may be removed and modified without warning.

## Copyright

Unless otherwise specificed, the copyright in this project are assigned as follows.

Copyright 2023 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.