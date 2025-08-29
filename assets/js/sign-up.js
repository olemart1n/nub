htmx.on("htmx:afterRequest", (evt) => {
  const notif = document.querySelector("#notification");
  if (notif && notif.dataset.error === "false") {
    setTimeout(() => {
      window.location.href = "/sign-in";
    }, 2000);
    return;
  }
  setTimeout(() => {
    notif.remove();
  }, 2000);
});

