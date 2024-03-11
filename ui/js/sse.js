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
    const duration = event.data.split(' ')[1]
    if (duration > .5 && duration < 1) {
      newElement.classList.remove('text-green-400')
      newElement.classList.add('orange')
    } else if (duration >= 1) {
      newElement.classList.remove('text-green-400')
      newElement.classList.add('red')
    }
    // Keep 500 messages in the container
    if (messagesContainer.childNodes.length > 499) {
      messagesContainer.removeChild(messagesContainer.firstChild)
    }
    messagesContainer.appendChild(newElement)
    messagesContainer.scrollTop = messagesContainer.scrollHeight
  }
})
