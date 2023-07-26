/*

https://portswigger.net/web-security/file-path-traversal/lab-validate-start-of-path

Lab: File path traversal, validation of start of path




	filename=/var/www/images/../../../etc/passwd


Step1: Use Burp Suite + FoxyProxy

Step2: intercept and modify a request that fetches a product image. 

Step3: Modify the filename parameter:

	/var/www/images/../../../etc/passwd

	
Result: Observe that the response contains the contents of the /etc/passwd file. 



*/
package main