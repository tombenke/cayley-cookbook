// Package foaf contains constants of the FOAF Concepts Vocabulary (RDF)
package foaf

import "github.com/cayleygraph/quad/voc"

func init() {
	voc.RegisterPrefix(Prefix, NS)
}

const (
	NS     = `http://xmlns.com/foaf/0.1/#`
	Prefix = `foaf:`
)

const (
	// Core

    //  An agent (eg. person, group, software or physical artifact).
    Agent = Prefix + `Agent`

    // A person.
    Person = Prefix + `Person`

    // A name for some thing.
    Name = Prefix + `name`

    // Title (Mr, Mrs, Ms, Dr. etc)
    Title = Prefix + `title`

    // An image that can be used to represent some thing (ie. those depictions which are particularly representative of something, eg. one's photo on a homepage).
    Img = Prefix + `img`

    // A thing depicted in this representation.
    Depiction = Prefix + `depiction`
    Depicts = Prefix + `depicts`

    // The family name of some person.
    FamilyName = Prefix + `familyName`

    // The first name of some person.
    GivenName = Prefix + `givenName`

    // The gender of this Agent (typically but not necessarily 'male' or 'female').
    Gender = Prefix + `gender`

    // A person known by this person (indicating some level of reciprocated interaction between the parties).
    Knows = Prefix + `knows`

    //  A location that something is based near, for some broadly human notion of near.
    BasedNear = Prefix + `based_near`

    // The age in years of some agent.
    Age = Prefix + `age`

    // The birthday of this Agent, represented in mm-dd string form, eg. '12-31'.
    Birthday = Prefix + `birthday`

    // Something that was made by this agent.
    Made = Prefix + `made`

    PrimaryTopic = Prefix + `primaryTopic (primaryTopicOf)`

    // A project (a collective endeavour of some kind).
    Project = Prefix + `Project`

    // An organization.
    Organization = Prefix + `Organization`

    // A class of Agents.
    Group = Prefix + `Group`

    // A string expressing what the user is happy for the general public (normally) to know about their current activity.
    Status = Prefix + `status`

    // Indicates a member of a Group
    Member = Prefix + `member`

    // A document.
    Document = Prefix + `Document`

    // An image.
    Image = Prefix + `Image`

    // SocialWeb

    // A short informal nickname characterising an agent (includes login identifiers, IRC and other chat nicknames).
    Nick = Prefix + `nick`

    // A personal mailbox, ie. an Internet mailbox associated with exactly one owner, the first owner of this mailbox. 
    Mbox = Prefix + `mbox`

    // A homepage for some thing.
    Homepage = Prefix + `homepage`

    // A weblog of some thing (whether person, group, company etc.).
    Weblog = Prefix + `weblog`

    // An OpenID for an Agent.
    Openid = Prefix + `openid`

    // A jabber ID for something.
    JabberID = Prefix + `jabberID`

    // sha1sum of a personal mailbox URI name - The sha1sum of the URI of an Internet mailbox associated with exactly one owner, the first owner of the mailbox.
    MboxSha1sum = Prefix + `mbox_sha1sum`

    // A page about a topic of interest to this person.
    Interest = Prefix + `interest`

    // A thing of interest to this person.
    TopicInterest = Prefix + `topic_interest`

    // A topic of some page or document.
    Topic = Prefix + `topic`

    // A workplace homepage of some person; the homepage of an organization they work for.
    WorkplaceHomepage = Prefix + `workplaceHomepage`

    // A work info homepage of some person; a page about their work for some organization.
    WorkInfoHomepage = Prefix + `workInfoHomepage`

    // A homepage of a school attended by the person.
    SchoolHomepage = Prefix + `schoolHomepage`

    // A link to the publications of this person.
    Publications = Prefix + `publications`

    // A current project this person works on.
    CurrentProject = Prefix + `currentProject`

    // A project this person has previously worked on.
    PastProject = Prefix + `pastProject`

    // Indicates an account held by this agent.
    Account = Prefix + `account`

    // An online account
    OnlineAccount = Prefix + `OnlineAccount`

    // Indicates the name (identifier) associated with this online account.
    AccountName = Prefix + `accountName`

    // Indicates a homepage of the service provide for this online account.
    AccountServiceHomepage = Prefix + `accountServiceHomepage`

    // A personal profile RDF document
    PersonalProfileDocument = Prefix + `PersonalProfileDocument`

    // A tipjar document for this agent, describing means for payment and reward.
    Tipjar = Prefix + `tipjar`

    // A sha1sum hash, in hex.
    Sha1 = Prefix + `sha1`

    // A derived thumbnail image.
    Thumbnail = Prefix + `thumbnail`

    // A logo representing some thing.
    Logo = Prefix + `logo`
)
