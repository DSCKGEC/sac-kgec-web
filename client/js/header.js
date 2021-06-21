const burger = document.querySelector(".burger");
const nav = document.querySelector(".navlinks");
const navlinks = document.querySelector(".navlinks li");
// Toggle nav
burger.addEventListener("click", () => {
    console.log('clicked');
    nav.classList.toggle("navactive");
});