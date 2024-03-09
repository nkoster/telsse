window.addEventListener('load', function () {
  const source = new EventSource('/sse')
  source.onmessage = function (event) {
    const messagesContainer = document.getElementById('messages')
    const newElement = document.createElement('div')
    newElement.textContent = event.data
    newElement.className = 'text-green-400 font-mono whitespace-pre-wrap m-0 p-0'
    messagesContainer.appendChild(newElement)

    // Eenvoudige manier om naar beneden te scrollen naar het nieuwste bericht
    messagesContainer.scrollTop = messagesContainer.scrollHeight
  }
})
