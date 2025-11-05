window.addEventListener("DOMContentLoaded", () => {
    const fileInputElement = document.querySelector('input[type="file"]');

    if (!window.FilePond) {
        console.warn("Filepond didn't load yet");
        return;
    }

    FilePond.registerPlugin(FilePondPluginImagePreview);
    const pond = FilePond.create(fileInputElement, {
        allowMultiple: true,
        instantUpload: false,
        imageResizeMode: "cover",
        stylePanelAspectRatio: "square",
        storeAsFile: true,
    });

    htmx.on("#postForm", "htmx:beforeRequest", () => {
        const loader = document.querySelector("#loader");
        loader.classList.remove("hidden");
        loader.classList.add("flex");
    });

    const checkBox = document.querySelector('input[type="checkbox"]');
    const uploadBtn = document.querySelector("form button");

    checkBox.addEventListener("change", () => {
        checkBox.checked
            ? uploadBtn.removeAttribute("disabled")
            : uploadBtn.setAttribute("disabled", true);
    });
});
