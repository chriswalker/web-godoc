let App = (() => {

  return {
    init: () => {
      // Attach event handler to input box for return handling
      input = document.getElementById('q')
      input.addEventListener('keydown', (e) => {
        if (e.keyCode == 13) {
          console.log('Return key pressed!')
          fetch('http://localhost:8000/search?q=' + input.value, { method: 'get' })
            .then(res => {
              return res.text()
            }).then(text => {
              console.log(text)
            })
          return false
        }
        return true
      })
    }
  }
})()

App.init()
