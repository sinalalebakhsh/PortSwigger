/*

https://portswigger.net/web-security/file-path-traversal/lab-validate-file-extension-null-byte-bypass


Lab: File path traversal, validation of file extension with null byte bypass



If an application requires that the user-supplied filename must end with an expected file extension, 
such as .png, 
then it might be possible to use a null byte to effectively terminate the file path before the required extension. 



For example:
	filename=../../../etc/passwd%00.png

Step1: Use Burp Suite + FoxyProxy

Step2: intercept and modify a request that fetches a product image. 

Step3: Modify the filename parameter:

	../../../etc/passwd%00.png


Result: Observe that the response contains the contents of the /etc/passwd file. 




































*/

package main