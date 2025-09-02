htmx.on("htmx:afterRequest", () => {
  const notif = document.querySelector("#notification");
  if (notif && notif.dataset.error === "true") {
    setTimeout(() => {
      console.log("notification should have been removed");
      notif.remove();
    }, 2000);
  }
});
