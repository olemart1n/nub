const menuToggler = document.querySelector("#menu-toggler");
const menu = document.querySelector("#menu");
menuToggler.addEventListener("click", () => {
  if (menu.classList.contains("translate-x-full")) {
    document.querySelector("header").classList.add("z-50");
    menu.classList.remove("translate-x-full");
  } else {
    menu.classList.add("translate-x-full");
    document.querySelector("header").classList.remove("z-50");
  }
});

