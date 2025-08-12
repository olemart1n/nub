const searchInput = document.querySelector(".search-input");
const dropdownContent = document.querySelector(".dropdown-content");
searchInput.addEventListener("input", function () {
  const searchTerm = this.value.toLowerCase();
  const options = document.querySelectorAll(".dropdown-content .option");

  options.forEach((option) => {
    const text = option.textContent.toLowerCase();
    if (text.includes(searchTerm)) {
      option.style.display = "block";
    } else {
      option.style.display = "none";
    }
  });
});

const options = document.querySelectorAll(".option");
options.forEach((option) => {
  option.addEventListener("click", function () {
    dropdownContent.style.display = "none";
    searchInput.value = option.innerHTML;
  });
});
searchInput.addEventListener("focus", () => {
  dropdownContent.style.display = "block";
});

document.addEventListener("click", (e) => {
  if (!dropdownContent.contains(e.target) && e.target !== searchInput) {
    dropdownContent.style.display = "none";
  }
});
