window.addEventListener('load', function () {
  const source = new EventSource('/sse')
  source.onmessage = function (event) {
    const messagesContainer = document.getElementById('messages')
    const newElement = document.createElement('div')
    newElement.textContent = event.data
    if (event.data === ':heartbeat') {
      return
    }
    newElement.className = 'text-green-400 font-mono whitespace-pre-wrap m-0 p-0'
    if (messagesContainer.childNodes.length > 500) {
      messagesContainer.removeChild(messagesContainer.firstChild)
    }
    messagesContainer.appendChild(newElement)
    messagesContainer.scrollTop = messagesContainer.scrollHeight
  }
})
