const slide = document.querySelector('.slide');
let position = 0;

setInterval(() => {
    slide.style.transform = `translateX(${position}px)`;
    position -= 700;
    if (position < -2100) {
        position = 0;
    }
}, 3000);


