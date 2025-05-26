// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        validateFormat()
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
      }, false)
    })
  })()
 
   if (document.getElementById('sangrenum').value === '0') {
    document.getElementById('sangre').disabled = true
  }

   
  if (document.getElementById('citologianum').value === '0') {
    document.getElementById('citologia').disabled = true
  }
   
  if (document.getElementById('yoganum').value === '0') {
    document.getElementById('yoga').disabled = true
  }
   
  if (document.getElementById('bailenum').value === '0') {
    document.getElementById('baile').disabled = true
  }
   
  if (document.getElementById('saludablenum').value === '0') {
    document.getElementById('saludable').disabled = true
  }