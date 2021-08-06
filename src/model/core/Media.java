package model.core;

import model.utility.*;

import java.util.ArrayList;
import java.util.List;

public class Media implements Playable {
    Title title;
    List<AudioTrack> audioTracks;
    List<VideoTrack> videoTracks;
    List<SubtitleTrack> subtitleTracks;
    List<Tag> tags;

    public Media(Title title) {
        this.title = title;
        this.audioTracks = new ArrayList<>();
        this.videoTracks = new ArrayList<>();
        this.subtitleTracks = new ArrayList<>();
        this.tags = new ArrayList<>();
    }

    public Snippet createSnippet(Title title, int start, int length) {
        return new Snippet(title, this, start, length);
    }

    public List<Track> getTracks() {
        List<Track> al = new ArrayList<>();
        al.addAll(this.audioTracks);
        al.addAll(this.videoTracks);
        al.addAll(this.subtitleTracks);
        return al;
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
