function toast_success(msg) {
  Toastify({
    text: msg,
    duration: 3000,
    close: true,
    gravity: "top",
    position: "center",
    style: {
      background: "linear-gradient(to right, #34d399, #10b981)",
    }
  }).showToast();
}

function toast_error(msg) {
  Toastify({
    text: msg,
    duration: 3000,
    close: true,
    gravity: "top",
    position: "center",
    style: {
      background: "linear-gradient(to right, #f87171, #ef4444)",
    }
  }).showToast();
}
