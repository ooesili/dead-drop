(function () {
  var fileInput = document.getElementById('file')
  var uploadInput = document.getElementById('upload')
  var state = {
    uploadDisabled: true
  }

  fileInput.onchange = onChangeFile
  renderView()

  function renderView () {
    uploadInput.disabled = state.uploadDisabled
  }

  function onChangeFile (e) {
    state.uploadDisabled = false
    renderView()
  }
})()
