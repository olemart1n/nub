htmx.on("htmx:afterRequest", (evt) => {
  setTimeout(() => {
    window.location.href = "/sign-up";
  }, 2000);
});