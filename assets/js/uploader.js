window.addEventListener("DOMContentLoaded", () => {
  const inputElement = document.querySelector('input[type="file"]');
  if (window.FilePond) {
    const pond = FilePond.create(inputElement);
    pond.setOptions({
      server: {
        process: (_, file, __, load, error, progress, abort) => {
          // Step 1: Ask our Go backend for upload details
          fetch("/sign-handler?filename=" + encodeURIComponent(file.name))
            .then((res) => res.json())
            .then((data) => {
              console.log(data);
              const xhr = new XMLHttpRequest();
              xhr.open("PUT", data.uploadURL, true);
              xhr.setRequestHeader("AccessKey", data.accessKey);

              xhr.upload.onprogress = (e) => {
                progress(e.lengthComputable, e.loaded, e.total);
              };

              xhr.onload = function () {
                if (xhr.status >= 200 && xhr.status < 300) {
                  load(data.publicURL); // Return the CDN URL
                } else {
                  error("Upload failed: " + xhr.statusText);
                }
              };

              xhr.onerror = () => error("Upload error");
              xhr.onabort = () => abort();

              xhr.send(file);
            });
        },
      },
    });
  } else {
    console.warn("FilePond didn’t load yet.");
  }
});
