
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import IndexPage from "./pages/Index";
import ProtectedRoute from "./components/protectedRoute";
import DashboardPage from "./pages/Dashboard";
import Login from "./pages/login/login";
import Signup from "./pages/signup/signup";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<IndexPage />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/dashboard" element={<ProtectedRoute><DashboardPage /></ProtectedRoute>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;



