import { useState, useEffect } from 'react';
import { Button } from "@/components/ui/button";

function LoginPage() {

    const [data, setData] = useState([]);

    const handleClick = () =>  {
        useEffect(() => {
            const fetchData = async () => {
                try {
                    const response = await fetch('localhost:3001/api/v1/ping');
                    const result = await response.json();
                    setData(result);
                }
                catch(error){
                    console.error('Error fetching data:', error)
                }
            }
            fetchData();
        })
        
    }

    return (
        <div>
            <h3>Auth0 Example</h3>
            <p>Zero friction identity infrastructure, built for developers</p>
            <Button onClick={handleClick}>Ping API</Button>
            <a href="localhost:3001/api/v1/ping">SignIn</a>
        </div>
    )
}

export default LoginPage


