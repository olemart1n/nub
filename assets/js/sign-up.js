htmx.on("htmx:afterRequest", (evt) => {
  setTimeout(() => {
    window.location.href = "/sign-in";
  }, 2000);
});