package model.meta;

import model.core.Media;
import model.core.Playable;
import model.utility.Celebrity;
import model.utility.Title;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Song implements Playable {
    Title title;
    List<Celebrity> artists;
    List<Celebrity> featuring;
    List<Media> variants;
    Map<Celebrity, Media> covers;
    Media cleanAudio;
    Media defaultVideo;

    public Song(Title title) {
        this.title = title;
        this.artists = new ArrayList<>();
        this.featuring = new ArrayList<>();
        this.variants = new ArrayList<>();
        this.covers = new HashMap<>();
    }

    public void addVariant(Media m) {
        this.variants.add(m);
    }

    public void addArtist(Celebrity c) {
        this.artists.add(c);
    }

    public void addCover(Celebrity c, Media m) {
        this.covers.put(c, m);
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        return null;
    }
}
