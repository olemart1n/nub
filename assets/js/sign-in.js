htmx.on("htmx:afterRequest", (evt) => {
  const notif = document.querySelector("#notification");
  if (notif && notif.dataset.error === "true") {
    setTimeout(() => {
      notif.remove();
    }, 2000);
  }
});
