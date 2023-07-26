/*
https://portswigger.net/web-security/file-path-traversal/lab-superfluous-url-decode

Lab: File path traversal, traversal sequences stripped with superfluous URL-decode

You can sometimes bypass this kind of sanitization by URL encoding, 
or even double URL encoding, 
the ../ characters, resulting in %2e%2e%2f or %252e%252e%252f respectively. 
Various non-standard encodings, such as ..%c0%af or ..%ef%bc%8f, may also do the trick.

For Burp Suite Professional users, 
Burp Intruder provides a predefined payload list (Fuzzing - path traversal), 
which contains a variety of encoded path traversal sequences that you can try. 


Step1: Use Burp Suite + FoxyProxy

Step2: intercept and modify a request that fetches a product image. 

Step3: Modify the filename parameter:
	..%252f..%252f..%252fetc/passwd

Result: Observe that the response contains the contents of the /etc/passwd file. 







































*/
package main