function animate(obj, target) {
    clearInterval(obj.timer);
    obj.timer = setInterval(function() {
        let step = (target - obj.offsetLeft) / 10;
        if (obj.offsetLeft >= target) {
            clearInterval(obj.timer);
        }
        obj.style.left = obj.offsetLeft + step + 'px';
    }, 10)
}

function jumpwindow(width, height, id) {
    const
}