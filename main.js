class Fps {
    constructor(interval, element) {
        this.lastTick = performance.now();
        this.lastNotify = this.lastTick;
        this.interval = interval;
        this.element = element;
        this.runningSum = 0;
        this.runningSamples = 0;
    }

    tick() {
        const now = performance.now();
        this.runningSum += (now - this.lastTick);
        this.runningSamples++;
        this.lastTick = now;

        if ((now - this.lastNotify) > this.interval) {
            this.notify(now);
        }

    }

    notify(now) {
        const avgFrame = this.runningSum / this.runningSamples;
        const fps = 1000 / avgFrame;
        this.element.innerText = `${fps.toFixed(2)}fps`;
        this.lastNotify = now;
        this.runningSamples = 0;
        this.runningSum = 0;
    }
}

const WIDTH = 800;
const HEIGHT = 800;

const elements = {
    canvas: {
        video: document.getElementById('video'),
        hidden: document.createElement('canvas'),
        visible: document.getElementById('canvas'),
        capture: document.getElementById('capture'),
        template: document.getElementById('template'),
        ctx: {}
    },
    loading: document.getElementById('loading'),
    fps: document.getElementById('fps'),
    options: [
        document.getElementById('grey_on'),
        document.getElementById('invert_on'),
        document.getElementById('noise_on'),
        document.getElementById('red_on'),
    ]
};

// Hidden canvas for rendering video directly onto

elements.canvas.hidden.width = WIDTH;
elements.canvas.hidden.height = HEIGHT;
elements.canvas.ctx.hidden = elements.canvas.hidden.getContext('2d');

// Target canvas for rendering effects to

elements.canvas.ctx.visible = elements.canvas.visible.getContext('2d');

// Setup target canvas for capture

elements.canvas.ctx.capture = elements.canvas.capture.getContext('2d');
elements.canvas.ctx.capture.drawImage(elements.canvas.template, 0, 0);


const getOptions = () => {

    let options = -1;
    for (let [i, el] of elements.options.entries()) {
        if (el.checked) {
            options = i
        }
    }

    return options;
};

// Start to capture webcam

navigator.mediaDevices.getUserMedia({ video: true, audio: false })
    .then(function (stream) {

        elements.canvas.video.srcObject = stream;
    })
    .catch(function (err) {

        throw err;
    });


