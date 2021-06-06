// 任务
export class TaskEvent{
    private name:Symbol
    private task:((...args:any) => {})
    private delay:number
    constructor(name:any, fn:((...args:any) => {}), delay = 0){
        this.name = Symbol(name);
        this.delay = delay;
        this.task = fn;
    }
    // 执行
    excute = () => {
        return new Promise(r =>
            setTimeout(() => {
                this.task && this.task()
            }, this.delay)
        )
    }
}

// 队列
export class EventsQueue {
    private queue:Array<TaskEvent>
    private isRunning:boolean
    constructor(){
        this.queue = [];
        this.isRunning = false;
    }

    // 添加任务
    addEvents = (e:TaskEvent) => {
        this.queue.push(e)
    }

    // 执行任务
    fire = async () => {
        if(this.isRunning)return
        while(this.queue.length !== 0){
            this.isRunning = true
            const evt = this.queue.splice(0, 1)[0];
            if(!evt)break
            await evt.excute()
        }
        this.isRunning = false
    }
}


const queue = new EventsQueue();
let uid = 0;
export const addEvents = (fn:(()=>{})) => {
    const task = new TaskEvent(uid, fn)
    queue.addEvents(task)
    uid++
}

export const fireTask = () => {
    queue.fire()
}


