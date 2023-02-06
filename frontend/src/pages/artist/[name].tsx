import { Artist } from "../../artist"
import { GetStaticPaths, GetStaticProps, NextPage } from "next"
import Image from "next/image"
import Head from "next/head"
import { getArtistByName, getArtistsLocal } from "../../fetch/server-side"

const ArtistPage: NextPage<{ artist: Artist }> = ({ artist }) => {
  // TODO: add client-side data fetch
  if (artist == undefined) {
    return <div>Error 500</div> // TODO: remove after client-side fetch implementation
  }
  return (
    <>
      <Head>
        <title>{`${artist.name} - GROUPIE TRACKER`}</title>
        <meta property={"title"} content={`${artist.name} - GROUPIE TRACKER`} key={"title"} />
        <meta
          name={"description"}
          content={`Some information about ${artist.name} and their concerts`}
        />
        <meta property={"og:title"} content={`${artist.name} - GROUPIE TRACKER`} key={"title"} />
        <meta
          name={"og:description"}
          content={`Some information about ${artist.name} and their concerts`}
        />
        {/* TODO: remove hardcoded og:image host */}
        <meta name={"og:image"} content={`https://groupie.mer.pw/${artist.image}`} />
      </Head>
      <div className={"w-fit m-auto text-center sm:text-left"}>
        <div className={"flex gap-7 flex-wrap justify-center"}>
          <div className={"m-auto sm:m-0 sm:text-right"}>
            <h1 className={"text-3xl font-bold mb-2"}>{artist.name}</h1>
            <hr />
            <h2 className={"text-xl my-2"}>Members:</h2>
            <ul className={"ml-1 text-xl font-light"}>
              {artist.members.map((member, key) => (
                <li key={key}>{member}</li>
              ))}
            </ul>
          </div>
          <div className={"order-first sm:order-none basis-full sm:basis-auto"}>
            <Image
              src={process.env.NEXT_PUBLIC_GROUPIE_BACKEND_HOST + artist.image}
              alt={`Image of ${artist.name}`}
              width={240}
              height={240}
              priority={true}
              className={"m-auto"}
            />
            <div className={"mt-2 m-auto"}>
              <h2 className={"text-xl font-light"}>Creation date: {artist.creationDate}</h2>
              <h2 className={"text-xl font-light"}>
                First album: {artist.firstAlbum.replaceAll("-", ".")}
              </h2>
            </div>
          </div>
          <div className={"text-center md:text-left"}>
            <h2 className={"text-xl mb-2"}>{"Concerts: "}</h2>
            <div className={"sm:flex sm:flex-col sm:flex-wrap"}>
              {artist.cities.map((city, key) => {
                return (
                  <div key={key} className={"text-xl font-light mb-2 max-w-screen-sm"}>
                    {`${city.name}: `}
                    <p className={"font-bold"}>
                      {city.dates
                        .map((date) => {
                          const [y, m, d] = date.split("-")
                          return `${d}.${m}.${y}`
                        })
                        .join(", ")}
                    </p>
                  </div>
                )
              })}
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export const getStaticPaths: GetStaticPaths<{ name: string }> = async () => {
  const artists: Artist[] = await getArtistsLocal()
  return {
    paths: artists.map((artist) => {
      return { params: { name: artist.name || "404" } }
    }),
    fallback: "blocking", // fallback tries to regenerate ArtistPage if Artist did not exist during building
  }
}

export const getStaticProps: GetStaticProps<{ artist: Artist }, { name: string }> = async (
  context
) => {
  if (context.params == undefined) {
    return { notFound: true, revalidate: 60 }
  }
  const artist = await getArtistByName(context.params.name)
  return artist ? { props: { artist: artist }, revalidate: 60 } : { notFound: true, revalidate: 60 }
}

export default ArtistPage
