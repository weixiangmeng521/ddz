import Home from "../views/home/home.jsx";
import Game from "../views/game/game.jsx";
import Test from "../views/cards/test.jsx";

const routes = [
    {
        path: "/",
        component: Home,
    },
    {
        path: "/game/:name",
        component: Game,
    },
    {
        path: "/cards/test",
        component: Test,
    },
];
export default routes;