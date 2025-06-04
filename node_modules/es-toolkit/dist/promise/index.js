'use strict';

Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

const delay = require('../_chunk/delay-_VMfFa.js');
const error_index = require('../error/index.js');

class Semaphore {
    capacity;
    available;
    deferredTasks = [];
    constructor(capacity) {
        this.capacity = capacity;
        this.available = capacity;
    }
    async acquire() {
        if (this.available > 0) {
            this.available--;
            return;
        }
        return new Promise(resolve => {
            this.deferredTasks.push(resolve);
        });
    }
    release() {
        const deferredTask = this.deferredTasks.shift();
        if (deferredTask != null) {
            deferredTask();
            return;
        }
        if (this.available < this.capacity) {
            this.available++;
        }
    }
}

class Mutex {
    semaphore = new Semaphore(1);
    get isLocked() {
        return this.semaphore.available === 0;
    }
    async acquire() {
        return this.semaphore.acquire();
    }
    release() {
        this.semaphore.release();
    }
}

async function timeout(ms) {
    await delay.delay(ms);
    throw new error_index.TimeoutError();
}

async function withTimeout(run, ms) {
    return Promise.race([run(), timeout(ms)]);
}

exports.delay = delay.delay;
exports.Mutex = Mutex;
exports.Semaphore = Semaphore;
exports.timeout = timeout;
exports.withTimeout = withTimeout;
