window.addEventListener("DOMContentLoaded", () => {
  const inputElement = document.querySelector('input[type="file"]');

  if (!window.FilePond) {
    console.warn("Filepond didn't load yet");
    return;
  }

  FilePond.registerPlugin(FilePondPluginImagePreview);
  const pond = FilePond.create(inputElement, {
    allowMultiple: true,
    instantUpload: false,
    imageResizeMode: "cover",
    stylePanelAspectRatio: "square",
     storeAsFile: true
  });


  htmx.on("#postForm", "htmx:beforeRequest", () => {
    console.log("htmx triggered")
    const loader = document.querySelector("#loader");
    loader.classList.remove("hidden");
    loader.classList.add("flex");
  });

});
