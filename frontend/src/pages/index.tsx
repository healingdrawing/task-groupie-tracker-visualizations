import { GetStaticProps, NextPage } from "next"
import { Artist } from "../artist"
import Link from "next/link"
import Image from "next/image"
import { getArtistsLocal } from "../fetch/server-side"

import { SWRConfig, SWRConfiguration } from "swr"
import { useArtists } from "../fetch/client-side"
import Head from "next/head"

const Home: NextPage<SWRConfiguration> = ({ fallback }) => {
  return (
    <SWRConfig value={{ fallback }}>
      <Head>
        <title>GROUPIE-TRACKER</title>
        <meta name={"description"} content={"Find concerts of your favourite artists"} />
        <meta property={"og:title"} content={"GROUPIE-TRACKER"} key={"title"} />
        <meta name={"og:description"} content={"Find concerts of your favourite artists"} />
      </Head>
      <ArtistsList />
    </SWRConfig>
  )
}

const ArtistsList = () => {
  const { artists, isError, isLoading } = useArtists()
  if (isLoading) {
    return <div>Loading...</div> //TODO: placeholder
  }
  if (isError) {
    // TODO: remove on production
    console.log(`Live refresh failed: '${isError.message}. Have you disabled CORS?`)
  }
  return (
    <div>
      <div className={"flex flex-wrap gap-1"}>
        {artists?.map((artist, key) => {
          return ArtistCard(artist, key)
        })}
      </div>
    </div>
  )
}

const ArtistCard = (artist: Artist, key: number) => (
  <div className={"mx-auto my-2 text-center p-1 w-max"} key={key}>
    <Link href={`/artist/${artist.name}`}>
      <Image
        src={process.env.NEXT_PUBLIC_GROUPIE_BACKEND_HOST + artist.image}
        alt={`Image of ${artist.name}`}
        width={240}
        height={240}
        blurDataURL={"placeholder"} // TODO: add base64 placeholder
        placeholder={"blur"}
        className={"rounded-full hover:brightness-125"}
      />
      <p className={"text-xl"}>{artist.name}</p>
    </Link>
  </div>
)

export const getStaticProps: GetStaticProps<SWRConfiguration<Artist[]>> = async () => {
  const artists: Artist[] = await getArtistsLocal()

  return {
    props: {
      fallback: {
        "/artists/": artists,
      },
    },
    revalidate: 60,
  }
}

export default Home
