import { useState } from "react"
import { Navigate } from "react-router-dom"

const Login = (props) => {

    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [redirect, setRedirect] = useState(false)

    const submit = async (e) => {
        e.preventDefault()

        const response = await fetch('http://localhost:8000/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        })

        const user = await response.json()

        setRedirect(true)
        props.setName(user.name)
    }

    if(redirect) {
        return <Navigate to="/" />
    }

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
            <input type="email" className="form-control" placeholder="Email address" required 
                    onChange={e => setEmail(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Password" required 
                    onChange={e => setPassword(e.target.value)}
            />

            <button className="btn btn-primary w-100 py-2" type="submit">Sign in</button>
        </form>
    )
}

export default Login