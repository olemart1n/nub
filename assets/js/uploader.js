window.addEventListener("DOMContentLoaded", () => {
  const inputElement = document.querySelector('input[type="file"]');
  if (!window.FilePond) {
    console.warn("Filepond didn't load yet");
    return;
  }
  FilePond.registerPlugin(
    FilePondPluginImagePreview,
    //  FilePondPluginImageEdit,
  );
  const pond = FilePond.create(inputElement, {
    allowMultiple: true,
    instantUpload: false,
    // imageResizeTargetWidth: 260,
    // imageResizeTargetHeight: 260,
    imageResizeMode: "cover",
    stylePanelAspectRatio: "square",
    //    imagePreviewHeight: 170,
  });

  console.log(inputElement);
  const form = document.querySelector("#postForm");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    // Create FormData with normal inputs
    const formData = new FormData(form);

    // Add FilePond files (multiple)
    pond.getFiles().forEach((fileItem) => {
      formData.append("images", fileItem.file); // "images" will be the key on the server
    });

    const title = formData.get("title");
    const location = formData.get("location");

    const rawTags = formData.get("tags"); // "erg,fjord,hiking"
    const tagsArray = rawTags.split(",").map((tag) => tag.trim());
    formData.set("tags", JSON.stringify(tagsArray)); // ✅ overwrite with JSON array
    formData.append("title", title);
    formData.append("location", location);

    const res = await fetch("/create-post", {
      method: "POST",
      body: formData,
    });

    const json = await res.json();
    console.log(json);
  });
});
