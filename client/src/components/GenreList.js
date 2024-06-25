import axios from 'axios';
import React, { useState, useEffect } from 'react'

import { Link } from 'react-router-dom';

const GenreList = () => {
    const [genres, setGenres] = useState([]);
    const [loaded, setLoaded] = useState(false);
    const [errorMessage, setErrorMessage] = useState(null);

    const fetchGenres = async () => {
        try {
            const result = await axios(`http://localhost:8000/api/genre/`);
            await setGenres(result.data.data);
            setLoaded(true);
        } catch (err) {
            setErrorMessage(err.response.message);
        }
    }

    useEffect(() => {
      fetchGenres();
    }, []);

    return (
        <>
            { !loaded ? (
                (() => {
                    if (errorMessage) {
                        return(
                            <div className="row">
                                <p>Opss... {errorMessage}</p>
                            </div>
                        );
                    } else {
                        return(
                            <div className="row">
                                <p>Loading...</p>
                            </div>
                        );
                    }
                })()
            ) : (
                <div className="row">
                    {genres.map((genre, index) => (
                        <div className="col-sm-2 mb-3" key={index}>
                            <div className="card">
                                <div className="card-body text-center">
                                    <Link to={`/genres/${genre.id}/movies`}>{genre.genre_name}</Link>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </>
    )
}

export default GenreList