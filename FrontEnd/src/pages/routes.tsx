import { BrowserRouter, Routes, Route } from 'react-router-dom';
import App from '../app/App';
import Home from "./Home";
import HomeConnecte from "./HomeConecte";
import Register from "./Register";
import Login from "./Login";
import CreateAnnonce from "./CreateAnnouncement"

export default function Routess() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<App />} />
                <Route path="/home" element={<Home />} />
                <Route path="/homeConnecte" element={<HomeConnecte />} />
                <Route path="/register" element={<Register />} />
                <Route path="/login" element={<Login />} />
                <Route path="/CreateAnnonce" element={<CreateAnnonce />} />
            </Routes>
        </BrowserRouter>
    )
}