import io from 'socket.io-client';

const uri = 'http://localhost:9527';
const options = {
    autoConnect: true,
    transports: ['websocket'],
    path: "/ws/",
};

export const Socket = io(uri, options);


Socket.on('message', message => {

    console.log('new message');
    console.log(message);
});


Socket.on('connect', message => {
    console.log('socket connected');
});


Socket.on("disconnect", () => {
    console.log("disconnect");
})


Socket.on("connect_error", error => {
    console.log(error);
});