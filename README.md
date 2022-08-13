# Ulpia

## What is Ulpia?

Ulpia is an aggregator. It scans multiple rare book databases and returns the most relevant results.

It is named after the [Ulpian library](https://en.wikipedia.org/wiki/Ulpian_Library), a Roman library founded by the Emperor Trajan in AD 114 and  the major library in the Western World upon the destruction of the Library of Alexandria in the 3rd century.

## How is Ulpia made?

Ulpia is a serverless SRU aggregator with a "back-end" in Go (compiled to WebAssembly) and a front-end in JavaScript and Bootstrap CSS.

### Serverless

An application that sends out requests to different SRU servers is typically made in a server-client environment. As I didn't want to set up a separate server for this, I made the application *serverless* by performing the web requests in the browser. This involved the use of a CORS proxy.

### Go / WebAssembly

Thanks to the integration of Go and WebAssembly, I could use Go for the "back-end" of Ulpia, most notably parsing the XML responses from the different SRU servers, and yielding the results in HTML templates.

### JavaScript

The actual request is done in JavaScript, making full use of the language's asynchronicity and speed.
