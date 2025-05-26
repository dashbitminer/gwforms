var filtroInput = document.getElementById("filtroInput");
    filtroInput.addEventListener("keyup", filtrarTabla);

    function filtrarTabla() {
      var inputValor = filtroInput.value.toUpperCase();
      var tabla = document.getElementById("tabla");
      var filas = tabla.getElementsByTagName("tr");

      for (var i = 0; i < filas.length; i++) {
        if (i === 0) { // Omitir la primera fila (encabezados)
          continue;
        }

        var celdas = filas[i].getElementsByTagName("td");
        var mostrarFila = false;

        for (var j = 0; j < celdas.length; j++) {
          var valorCelda = celdas[j].textContent || celdas[j].innerText;
          if (valorCelda.toUpperCase().indexOf(inputValor) > -1) {
            mostrarFila = true;
            break;
          }
        }

        if (mostrarFila) {
          filas[i].style.display = "";
        } else {
          filas[i].style.display = "none";
        }
      }
    }

    function enviarFormulario(url) {
      var form = document.getElementById("formualrioRegistro");
      form.action = url;
      form.submit();
    }
    function confirmAction() {
      if (window.confirm("¿Esta seguro de eliminar el registro?")) {
          // User clicked "OK"
          enviarFormulario('/sanamente/modulo/eliminar');
      } else {
          // User clicked "Cancel";
      }
  }
   