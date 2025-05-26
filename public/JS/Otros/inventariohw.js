// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
  
        form.classList.add('was-validated')
        
      }, false)
  
    })
  })()

  $(document).ready(function() {
    $('#agregar-fila').click(function() {
      var html = '<tr>';
      html += '<td><input type="text" name="dispositivo[]" class="form-control"></td>';
      html += '<td><input type="text" name="modelo[]" class="form-control"></td>';
      html += '<td><input type="text" name="serie[]" class="form-control"></td>';
      html += '<td><input type="text" name="marca[]" class="form-control"></td>';
      html += '<td><input type="text" name="nActivo[]" class="form-control"></td>';
      html += '<td><input type="text" name="fechaRecepcion[]" class="form-control"></td>';
      html += '</tr>';
      $('table tbody').append(html);
    });
  });
