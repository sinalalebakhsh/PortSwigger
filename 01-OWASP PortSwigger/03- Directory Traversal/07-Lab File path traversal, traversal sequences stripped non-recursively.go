/*
https://portswigger.net/web-security/file-path-traversal/lab-sequences-stripped-non-recursively

Lab: File path traversal, traversal sequences stripped non-recursively


Step1: Use Burp Suite to intercept

Step2: modify a request that fetches a product image. 

Step3: Modify the filename parameter, giving it the value: 

	....//....//....//etc/passwd

Result: 
	Observe that the response contains the contents of the /etc/passwd file. 
	

























*/
package main