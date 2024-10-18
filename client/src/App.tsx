import './App.css'

import { Route, Routes } from 'react-router-dom'
import { NavBar } from './components/ui/navigation/desktop/nav-bar'
import { AuthenticationGuard } from "./components/authentication-guard";

import HomePage from './pages/landing-page'
import ProfilePage from './pages/profile-page'
import CallbackPage from './pages/callback-page';
import LandingPage from './pages/landing-page';

function App() {

  return (
    <>
      <Routes>
        <Route 
          path="/" 
          element={<LandingPage/>}
        /> {/* 👈 Renders at /app/ */}

        <Route 
          path="/callback" 
          element={<CallbackPage />} 
        />
        
        <Route 
          path="/home" 
          element={<AuthenticationGuard component={HomePage} />}
        />
        
        <Route
          path="/profile"
          element={<AuthenticationGuard component={ProfilePage} />}
        />
      </Routes>
      
    </>
  )
}

export default App
