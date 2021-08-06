package model;

import model.core.Media;
import model.core.Playlist;
import model.meta.*;
import model.utility.Celebrity;
import model.utility.Title;
import model.utility.VideoTrack;

public class Requirements {
    public static void main(String[] args) {
        testMovie();
        testSerie();
        testPlaylist();
        testMusic();
        testYoutube();
        testImages();
        testSnippets();
    }

    private static void testMovie() {
        // Verlinkung von Schauspielern
        // IMDB-Integration
        // Verlinkung mehrsprachiger Video & Audio Tracks
        // Verlinkung verschiedener Cuts
        Movie mov1 = new Movie(new Title("LOTR 2"));
        Media med1 = new Media(new Title("LOTR 2 - Cinema Cut"));
        mov1.addCut(med1);
        Media med2 = new Media(new Title("LOTR 2 - Directors Cut"));
        mov1.addCut(med2);
        // Filmreihen (e.g. Triologien
        Movie mov2 = new Movie(new Title("LOTR 3"));
        mov1.setSequel(mov2);
    }

    private static void testSerie() {
        // Verlinkung von Schauspielern
        // IMDB-Integration
        // Verlinkung mehrsprachiger Video & Audio Tracks
        // Episoden und Folgen
        Series ser1 = new Series(new Title("The Witcher"));
        Season sea1 = new Season(ser1, new Title("The Witcher - Season 1"));
        sea1.addEpisode(new Media(new Title("Folge #1")));
        sea1.addEpisode(new Media(new Title("Folge #2")));
        Season sea2 = new Season(ser1, new Title("The Witcher - Season 2"));
        sea2.addEpisode(new Media(new Title("Folge #3")));
        sea2.addEpisode(new Media(new Title("Folge #4")));
    }

    private static void testPlaylist() {
        // Mixed Playlisten (e.g. Youtube und Musik)
        Playlist pl = new Playlist(new Title("Test PL"));
        pl.addItem(new Media(new Title("Youtube Video")));
        pl.addItem(new Song(new Title("Spotify Song")));
        // Sowohl alg. Lied als auch Varianten (e.g. Coverversion) definierbar
    }

    private static void testMusic() {
        // Verlinkung mit original Titel (e.g. Spotify ID)
        // Varianten (e.g. Live, Studio, Lyric)
        Song s = new Song(new Title("Around the World"));
        s.addArtist(new Celebrity("Daft Punk"));
        s.addVariant(new Media(new Title("Around the World - Official Lyrics")));
        s.addVariant(new Media(new Title("Around the World - Official Video")));
        // Coverversion definierbar / differentierbar
        Celebrity c = new Celebrity("Max Power");
        Media m = new Media(new Title("Datf Punk on my Ukulele"));
        s.addCover(c, m);
        // Alben
        Album a = new Album(new Title("Best of Daft Punk"));
        a.addSong(s);
    }

    private static void testYoutube() {
        new Media(new Title("Pietsmiet Season #7 Trailer"));
    }

    private static void testImages() {
        ImageSet is = new ImageSet(new Title("Vacation in Paris"));
        is.addImage(new Media(new Title("IMG_1234.png")));
    }

    private static void testSnippets() {
        // aus Videos
        Media med1 = new Media(new Title("Full Concert"));
        Media med2 = med1.createSnippet(new Title("Song #3 from Full Concert"), 23, 60);
        // aus YT-Video als Musik (Variante)
        Song s = new Song(new Title("Whatever"));
        s.addVariant(med2);
        // aus Film als Musik
        Movie mov = new Movie(new Title("LOTR 2"));
        mov.addCut(new Media(new Title("LOTR 2 - Directors Cut")));
        Media med3 = mov.getCuts(0).createSnippet(new Title("Intro"), 0, 60);
        s.addVariant(med3);
    }
}
