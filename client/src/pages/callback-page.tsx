import { useAuth0 } from "@auth0/auth0-react";

export default function CallbackPage() {
    const { error } = useAuth0();

    if (error) {
        return (
            <p>Callback error</p>
        );
      }

    return (
        <p>Callback ok</p>
    )
}