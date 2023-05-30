// Reference: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
const xhr = new XMLHttpRequest()
xhr.open('GET', 'http://demo.local:8080/')
xhr.setRequestHeader('Content-Type', 'text/xml') // forces preflight
xhr.onreadystatechange = () => {
	const mainDiv = document.getElementById('main')
	mainDiv.textContent = xhr.responseText
}
xhr.send()
