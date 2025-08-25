    const tagInput = document.querySelector("#tagInput");
    const addedTags = document.querySelector("#addedTags");
    const predefinedTagsDiv = document.querySelector("#predefined-tags");

    let tags = []

    let predefinedTags = ["nubs", "box", "pyrimad", "damage-holes", "clamps", "bedrock", "bedrock", "pinch-holes", "temple", "filler-stone"]

    function renderTags() {
      addedTags.innerHTML = "";
      tags.forEach(tag => {
        const span = document.createElement("span");
        span.className = "tag";
        span.textContent = tag;
        span.onclick = () => removeTag(tag);
        addedTags.appendChild(span);
        });
    }

    function renderPredefinedTags() {
      predefinedTagsDiv.innerHTML = "";
      predefinedTags.forEach(tag => {
        const span = document.createElement("span");
        span.className = "outline p-0.5 rounded text-slate-100 outline-slate-200 cursor-pointer";
        span.textContent = tag;
        
        span.addEventListener("click", () => {
            addTag(tag);
            removePredefinedTag(tag)
        });
        predefinedTagsDiv.appendChild(span);
        });
    }

    function addTag(tag) {
      tag = tag.trim();
      if (tag && !tags.includes(tag)) {
        tags.push(tag);
        renderTags();
      }
    }

    function removeTag(tag) {
      tags = tags.filter(t => t !== tag);
      renderTags();
    }
    function removePredefinedTag(tag) {
      predefinedTags = predefinedTags.filter(t => t !== tag);
      renderPredefinedTags();
    }

    // Input: add tag on Enter
    tagInput.addEventListener("keydown", e => {
      if (e.key === "Enter") {
        e.preventDefault();
        addTag(tagInput.value);
        tagInput.value = "";
      }
    });

    renderPredefinedTags()







// ADD THIS TO THE HTMX CODE IN uploader.js if needed.
    // const hiddenInput = form.querySelector("input[name='tags']");
    // hiddenInput.value = JSON.stringify(tags);