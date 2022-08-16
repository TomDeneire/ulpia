package ujson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var googlebooksJSON = []byte(`
{
    "kind": "books#volumes",
    "totalItems": 872,
    "items": [
      {
        "kind": "books#volume",
        "id": "8WCNT4LSkFcC",
        "etag": "hN9xY/55Zj8",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/8WCNT4LSkFcC",
        "volumeInfo": {
          "title": "Justus Lipsius (1547-1606)",
          "subtitle": "een geleerde en zijn europese netwerk : catalogus van de tentoonstellung in de Centrale Bibliotheek te Leuven, 18 oktober-20 december 2006",
          "authors": [
            "Jeanine Landtsheer",
            "Dirk Sacré",
            "Chris Coppens"
          ],
          "publisher": "Leuven University Press",
          "publishedDate": "2006",
          "industryIdentifiers": [
            {
              "type": "ISBN_10",
              "identifier": "905867567X"
            },
            {
              "type": "ISBN_13",
              "identifier": "9789058675675"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 626,
          "printType": "BOOK",
          "categories": [
            "Humanists"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "3.5.2.0.preview.1",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=8WCNT4LSkFcC&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=8WCNT4LSkFcC&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "nl",
          "previewLink": "http://books.google.be/books?id=8WCNT4LSkFcC&pg=PA297&dq=lipsius+admiranda+1598&hl=&cd=1&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=8WCNT4LSkFcC&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Justus_Lipsius_1547_1606.html?hl=&id=8WCNT4LSkFcC"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": false
          },
          "webReaderLink": "http://play.google.com/books/reader?id=8WCNT4LSkFcC&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "... the Visualization of Ancient Rome &#39; , Sixteenth Century Journal , 35 ( 2004 ) , 97-131 ; K. Enenkel , &#39; Ein Plädoyer für den Imperialismus : Justus \u003cb\u003eLipsius\u003c/b\u003e &#39; kulturhistorische Monographie \u003cb\u003eAdmiranda\u003c/b\u003e sive De Magnitudine Romana ( \u003cb\u003e1598\u003c/b\u003e )&nbsp;..."
        }
      },
      {
        "kind": "books#volume",
        "id": "SbPjBQAAQBAJ",
        "etag": "pNwEsBNYAi8",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/SbPjBQAAQBAJ",
        "volumeInfo": {
          "title": "Jan Moretus and the Continuation of the Plantin Press (2 Vols.)",
          "subtitle": "A Bibliography of the Works published and printed by Jan Moretus I in Antwerp (1589-1610)",
          "authors": [
            "Dirk Imhof"
          ],
          "publisher": "Hotei Publishing",
          "publishedDate": "2014-10-14",
          "description": "This extensive bibliography contains 704 descriptions of the Jan Moretus editions and lists over 500 announcements that he printed for the city of Antwerp.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9789004287884"
            },
            {
              "type": "ISBN_10",
              "identifier": "9004287884"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 1124,
          "printType": "BOOK",
          "categories": [
            "History"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "preview-1.0.0",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=SbPjBQAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=SbPjBQAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=SbPjBQAAQBAJ&pg=PA314&dq=lipsius+admiranda+1598&hl=&cd=2&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=SbPjBQAAQBAJ&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Jan_Moretus_and_the_Continuation_of_the.html?hl=&id=SbPjBQAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": false
          },
          "webReaderLink": "http://play.google.com/books/reader?id=SbPjBQAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "Lipsius (1967), no. 46) and Arch. 75, fols, 43v and 49r, and Arch. 21, fol. 335 lft). On the Admiranda, see Tom Deneire, &quot;Justus \u003cb\u003eLipsius&#39;s Admiranda\u003c/b\u003e (\u003cb\u003e1598\u003c/b\u003e) and the Officina Plantiniana: mixing Otium with&nbsp;..."
        }
      },
      {
        "kind": "books#volume",
        "id": "NS9yEAAAQBAJ",
        "etag": "Pna+5NsdqUQ",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/NS9yEAAAQBAJ",
        "volumeInfo": {
          "title": "Justus Lipsius, Monita et exempla politica / Political Admonitions and Examples",
          "authors": [
            "Jan Papy",
            "Toon Van Houdt",
            "Marijke Janssens"
          ],
          "publisher": "Leuven University Press",
          "publishedDate": "2022-05-30",
          "description": "In 17th-century intellectual life, the ideas of the Renaissance humanist Justus Lipsius (1547–1606) were omnipresent. The publication of his Politica in 1589 had made Lipsius' name as an original and controversial political thinker. The sequel, the Monita et exempla politica (Political admonitions and examples), published in 1605, was meant as an illustration of Lipsius political thought as expounded in the Politica. Its aim was to offer concrete models of behavior for rulers against the background of Habsburg politics. Lipsius' later political treatise also forms an indispensable key to interpret the place and function of the Politica in Lipsius’ political discourse and in early modern political thought. The Political admonitions and examples – widely read, edited, and translated in the 17th and 18th centuries – show Lipsius’ pivotal role in the genesis of modern political philosophy.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9789462703056"
            },
            {
              "type": "ISBN_10",
              "identifier": "9462703051"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 696,
          "printType": "BOOK",
          "categories": [
            "History"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "preview-1.0.0",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=NS9yEAAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=NS9yEAAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=NS9yEAAAQBAJ&pg=PA44&dq=lipsius+admiranda+1598&hl=&cd=3&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=NS9yEAAAQBAJ&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Justus_Lipsius_Monita_et_exempla_politic.html?hl=&id=NS9yEAAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": false
          },
          "webReaderLink": "http://play.google.com/books/reader?id=NS9yEAAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "Some of these elements were already present in \u003cb\u003eLipsius\u003c/b\u003e&#39; \u003cb\u003eAdmiranda\u003c/b\u003e (\u003cb\u003e1598\u003c/b\u003e) and Panegyricus (1600). Thus, from the outset of the Panegyricus, Lipsius takes Pliny&#39;s preliminary prayer as a starting point to stress the divine foundation of&nbsp;..."
        }
      },
      {
        "kind": "books#volume",
        "id": "AAVPWfmuNnsC",
        "etag": "iXW//AFF7ak",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/AAVPWfmuNnsC",
        "volumeInfo": {
          "title": "(Un)masking the Realities of Power",
          "subtitle": "Justus Lipsius and the Dynamics of Political Writing in Early Modern Europe",
          "authors": [
            "Erik Bom",
            "Marijke Janssens",
            "Toon van Houdt"
          ],
          "publisher": "BRILL",
          "publishedDate": "2010-12-10",
          "description": "Starting from Justus Lipsius's Monita et exempla politica (1605), this book offers a collection of essays dealing with the disputed Macchiavellian, Tacitean or Neostoic character of Lipsius's political thought, and its impact on the dynamics of political discourse in Early Modern Europe.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9789004191280"
            },
            {
              "type": "ISBN_10",
              "identifier": "9004191283"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 347,
          "printType": "BOOK",
          "categories": [
            "Social Science"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "1.2.1.0.preview.1",
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=AAVPWfmuNnsC&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=AAVPWfmuNnsC&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=AAVPWfmuNnsC&pg=PA67&dq=lipsius+admiranda+1598&hl=&cd=4&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=AAVPWfmuNnsC&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Un_masking_the_Realities_of_Power.html?hl=&id=AAVPWfmuNnsC"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": false
          },
          "webReaderLink": "http://play.google.com/books/reader?id=AAVPWfmuNnsC&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "Justus Lipsius and the Dynamics of Political Writing in Early Modern Europe Erik Bom, Marijke Janssens, Toon van Houdt ... 76 Cf. e.g. Deneire T., “Justus \u003cb\u003eLipsius&#39;s Admiranda\u003c/b\u003e (\u003cb\u003e1598\u003c/b\u003e) and history and exemplarity in the work of lipsius 67."
        }
      },
      {
        "kind": "books#volume",
        "id": "pzMrDwAAQBAJ",
        "etag": "mq1eIhV8nms",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/pzMrDwAAQBAJ",
        "volumeInfo": {
          "title": "The Aura of the Word in the Early Age of Print (1450-1600)",
          "authors": [
            "Samuel Mareel"
          ],
          "publisher": "Routledge",
          "publishedDate": "2017-07-05",
          "description": "Did the invention of movable type change the way that the word was perceived in the early modern period? In his groundbreaking essay \"The Work of Art in the Age of Mechanical Reproduction,\" the cultural critic Walter Benjamin argued that reproduction drains the image of its aura, by which he means the authority that a work of art obtains from its singularity and its embeddedness in a particular context. The central question in The Aura of the Word in the Early Age of Print (1450-1600) is whether the dissemination of text through print had a similar effect on the status of the word in the early modern period. In this volume, contributors from a variety of fields look at manifestations of the early modern word (in English, French, Latin, Dutch, German and Yiddish) as entities whose significance derived not simply from their semantic meaning but also from their relationship to their material support, to the physical context in which they are located and to the act of writing itself. Rather than viewing printed text as functional and lacking in materiality, contributors focus on how the placement of a text could affect its meaning and significance. The essays also consider the continued vitality of pre-printing-press kinds of text such as the illuminated manuscript; and how new practices, such as the veneration of handwriting, sprung up in the wake of the invention of movable type.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9781351546102"
            },
            {
              "type": "ISBN_10",
              "identifier": "1351546104"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 232,
          "printType": "BOOK",
          "categories": [
            "Antiques & Collectibles"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "0.2.0.0.preview.1",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=pzMrDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=pzMrDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=pzMrDwAAQBAJ&pg=PA128&dq=lipsius+admiranda+1598&hl=&cd=5&source=gbs_api",
          "infoLink": "https://play.google.com/store/books/details?id=pzMrDwAAQBAJ&source=gbs_api",
          "canonicalVolumeLink": "https://play.google.com/store/books/details?id=pzMrDwAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "FOR_SALE",
          "isEbook": true,
          "listPrice": {
            "amount": 39.21,
            "currencyCode": "EUR"
          },
          "retailPrice": {
            "amount": 28.89,
            "currencyCode": "EUR"
          },
          "buyLink": "https://play.google.com/store/books/details?id=pzMrDwAAQBAJ&rdid=book-pzMrDwAAQBAJ&rdot=1&source=gbs_api",
          "offers": [
            {
              "finskyOfferType": 1,
              "listPrice": {
                "amountInMicros": 39210000,
                "currencyCode": "EUR"
              },
              "retailPrice": {
                "amountInMicros": 28890000,
                "currencyCode": "EUR"
              }
            }
          ]
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": true,
            "acsTokenLink": "http://books.google.be/books/download/The_Aura_of_the_Word_in_the_Early_Age_of-sample-pdf.acsm?id=pzMrDwAAQBAJ&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
          },
          "webReaderLink": "http://play.google.com/books/reader?id=pzMrDwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "For more information on Lipsius&#39;s Admiranda, sive de magnitudine Romana libri quattuor (Antwerp: J. Moretus, 1598), see Tom Deneire, “Justus \u003cb\u003eLipsius&#39;s Admiranda\u003c/b\u003e (\u003cb\u003e1598\u003c/b\u003e) and the Officina Plantiniana: mixing otium with negotium,” in Iam&nbsp;..."
        }
      },
      {
        "kind": "books#volume",
        "id": "tCQxDwAAQBAJ",
        "etag": "M7+uShIKyhs",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/tCQxDwAAQBAJ",
        "volumeInfo": {
          "title": "The Aura of the Word in the Early Age of Print (1450?600)",
          "authors": [
            "Samuel Mareel"
          ],
          "publisher": "Routledge",
          "publishedDate": "2017-07-05",
          "description": "Did the invention of movable type change the way that the word was perceived in the early modern period? In his groundbreaking essay \"The Work of Art in the Age of Mechanical Reproduction,\" the cultural critic Walter Benjamin argued that reproduction drains the image of its aura, by which he means the authority that a work of art obtains from its singularity and its embeddedness in a particular context. The central question in The Aura of the Word in the Early Age of Print (1450-1600) is whether the dissemination of text through print had a similar effect on the status of the word in the early modern period. In this volume, contributors from a variety of fields look at manifestations of the early modern word (in English, French, Latin, Dutch, German and Yiddish) as entities whose significance derived not simply from their semantic meaning but also from their relationship to their material support, to the physical context in which they are located and to the act of writing itself. Rather than viewing printed text as functional and lacking in materiality, contributors focus on how the placement of a text could affect its meaning and significance. The essays also consider the continued vitality of pre-printing-press kinds of text such as the illuminated manuscript; and how new practices, such as the veneration of handwriting, sprung up in the wake of the invention of movable type.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9781351546096"
            },
            {
              "type": "ISBN_10",
              "identifier": "1351546090"
            }
          ],
          "readingModes": {
            "text": true,
            "image": true
          },
          "pageCount": 232,
          "printType": "BOOK",
          "categories": [
            "Antiques & Collectibles"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "1.1.1.0.preview.3",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=tCQxDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=tCQxDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=tCQxDwAAQBAJ&pg=PT155&dq=lipsius+admiranda+1598&hl=&cd=6&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=tCQxDwAAQBAJ&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/The_Aura_of_the_Word_in_the_Early_Age_of.html?hl=&id=tCQxDwAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": true,
            "acsTokenLink": "http://books.google.be/books/download/The_Aura_of_the_Word_in_the_Early_Age_of-sample-epub.acsm?id=tCQxDwAAQBAJ&format=epub&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
          },
          "pdf": {
            "isAvailable": true,
            "acsTokenLink": "http://books.google.be/books/download/The_Aura_of_the_Word_in_the_Early_Age_of-sample-pdf.acsm?id=tCQxDwAAQBAJ&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
          },
          "webReaderLink": "http://play.google.com/books/reader?id=tCQxDwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "... on Lipsius&#39;s Admiranda, sive de magnitudine Romana libri quattuor (Antwerp: J. Moretus, 1598), see Tom Deneire, “Justus \u003cb\u003eLipsius&#39;s Admiranda\u003c/b\u003e (\u003cb\u003e1598\u003c/b\u003e) and the Officina Plantiniana: mixing otium with negotium,” in Iam illustravit omnia."
        }
      },
      {
        "kind": "books#volume",
        "id": "prc5d99Rd5kC",
        "etag": "4zQs9LEh0vc",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/prc5d99Rd5kC",
        "volumeInfo": {
          "title": "Lipsius en Leuven",
          "subtitle": "catalogus van de tentoonstelling in de Centrale Bibliotheek te Leuven, 18 september-17 oktober 1997",
          "authors": [
            "Gilbert Tournoy",
            "Jan Papy",
            "Jeanine Landtsheer",
            "Katholieke Universiteit te Leuven (1970- ). Universiteitsbibliotheek"
          ],
          "publisher": "Leuven University Press",
          "publishedDate": "1997",
          "description": "Justus Lipsius (Overijse, 18 oktober 1547 - Leuven, 23 maart 1606) was een Zuid-Nederlandse humanist, filoloog en historiograaf.Lipsius studeerde in het katholieke Leuven en doceerde in het lutherse Jena en in het calvinistische Leiden. Uiteindelijk koos hij tegen het geweld van de Tachtigjarige Oorlog en zou hij het voortduren van de Spaanse heerschappij over de Nederlanden verkieslijker achten dan langdurige oorlog en opstand. Zo eindigde hij weer als docent in Leuven aan de \"katholieke\" kant. Zijn voornaamste werk is De Constantia (1584).Tentoonstelling in de Centrale Bibliotheek te Leuven, van 18 september tot 17 oktober 1997.",
          "industryIdentifiers": [
            {
              "type": "ISBN_10",
              "identifier": "906186836X"
            },
            {
              "type": "ISBN_13",
              "identifier": "9789061868361"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 387,
          "printType": "BOOK",
          "categories": [
            "Electronic book"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "3.3.5.0.preview.1",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=prc5d99Rd5kC&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=prc5d99Rd5kC&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "nl",
          "previewLink": "http://books.google.be/books?id=prc5d99Rd5kC&pg=PA335&dq=lipsius+admiranda+1598&hl=&cd=7&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=prc5d99Rd5kC&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Lipsius_en_Leuven.html?hl=&id=prc5d99Rd5kC"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "ALL_PAGES",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": true,
            "acsTokenLink": "http://books.google.be/books/download/Lipsius_en_Leuven-sample-pdf.acsm?id=prc5d99Rd5kC&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
          },
          "webReaderLink": "http://play.google.com/books/reader?id=prc5d99Rd5kC&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "III Portretmedaille van \u003cb\u003eLipsius\u003c/b\u003e Medailles in goud en zilver , 44 mm , hangoog , \u003cb\u003e1598\u003c/b\u003e of 1601 ( ? ) ... Wel had \u003cb\u003eLipsius\u003c/b\u003e in maart \u003cb\u003e1598\u003c/b\u003e zijn \u003cb\u003eAdmiranda\u003c/b\u003e aan aartshertog Albrecht opgedragen ( zie cat . nr . 27 ) , zodat hij misschien \u003cb\u003eLipsius\u003c/b\u003e&nbsp;..."
        }
      },
      {
        "kind": "books#volume",
        "id": "ijE9DwAAQBAJ",
        "etag": "Ji/hrDqJHQs",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/ijE9DwAAQBAJ",
        "volumeInfo": {
          "title": "Ancient Models in the Early Modern Republican Imagination",
          "publisher": "BRILL",
          "publishedDate": "2017-10-20",
          "description": "Ancient Models in the Early Modern Republican Imagination offers a new approach to the study of the classical dimensions of early modern republican thought by analysing its specific and concrete uses of ancient republican models.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9789004351387"
            },
            {
              "type": "ISBN_10",
              "identifier": "9004351388"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 352,
          "printType": "BOOK",
          "categories": [
            "History"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "preview-1.0.0",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=ijE9DwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=ijE9DwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=ijE9DwAAQBAJ&pg=PA72&dq=lipsius+admiranda+1598&hl=&cd=8&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=ijE9DwAAQBAJ&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Ancient_Models_in_the_Early_Modern_Repub.html?hl=&id=ijE9DwAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": false
          },
          "webReaderLink": "http://play.google.com/books/reader?id=ijE9DwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "“He that will attain to glory or empire, let him turn to the ancient discipline”, \u003cb\u003eLipsius\u003c/b\u003e exclaimed.32 This assertion clearly attested to the ... 35 Justus \u003cb\u003eLipsius\u003c/b\u003e, \u003cb\u003eAdmiranda\u003c/b\u003e, sive de magnitudine romana (Antwerp, \u003cb\u003e1598\u003c/b\u003e), Weststeijn 72."
        }
      },
      {
        "kind": "books#volume",
        "id": "J-EqDwAAQBAJ",
        "etag": "JlzSzAqmZBE",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/J-EqDwAAQBAJ",
        "volumeInfo": {
          "title": "Erudite Eyes",
          "subtitle": "Friendship, Art and Erudition in the Network of Abraham Ortelius (1527-1598)",
          "authors": [
            "Tine Luk Meganck"
          ],
          "publisher": "BRILL",
          "publishedDate": "2017-06-26",
          "description": "Erudite Eyes explores how friendship between artists and humanists in the network of Abraham Ortelius (1527-1598) produced an antiquarian culture that yielded new knowledge on local antiquities and distant civilizations and that articulated artistic practice between Bruegel and Rubens.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9789004342484"
            },
            {
              "type": "ISBN_10",
              "identifier": "9004342486"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 346,
          "printType": "BOOK",
          "categories": [
            "Art"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "preview-1.0.0",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=J-EqDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=J-EqDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=J-EqDwAAQBAJ&pg=PA300&dq=lipsius+admiranda+1598&hl=&cd=9&source=gbs_api",
          "infoLink": "http://books.google.be/books?id=J-EqDwAAQBAJ&dq=lipsius+admiranda+1598&hl=&source=gbs_api",
          "canonicalVolumeLink": "https://books.google.com/books/about/Erudite_Eyes.html?hl=&id=J-EqDwAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "NOT_FOR_SALE",
          "isEbook": false
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": false
          },
          "webReaderLink": "http://play.google.com/books/reader?id=J-EqDwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "\u003cb\u003eLipsius 1598\u003c/b\u003e J. \u003cb\u003eLipsius\u003c/b\u003e, \u003cb\u003eAdmiranda\u003c/b\u003e sive de magnitudine romana libri IV, Antwerp \u003cb\u003e1598\u003c/b\u003e. \u003cb\u003eLipsius\u003c/b\u003e 1609 J. \u003cb\u003eLipsius\u003c/b\u003e, De Vesta et Vestablibus syntagma, Antwerp 1609. \u003cb\u003eLipsius\u003c/b\u003e 1978 J. \u003cb\u003eLipsius\u003c/b\u003e, Iusti Lipsi Epistolae (A. Gerlo, M.A. Nauwelaerts,&nbsp;..."
        }
      },
      {
        "kind": "books#volume",
        "id": "DSY-DwAAQBAJ",
        "etag": "MjcHmVpNfF8",
        "selfLink": "https://www.googleapis.com/books/v1/volumes/DSY-DwAAQBAJ",
        "volumeInfo": {
          "title": "The End of Fortuna and the Rise of Modernity",
          "authors": [
            "Arndt Brendecke",
            "Peter Vogt"
          ],
          "publisher": "Walter de Gruyter GmbH & Co KG",
          "publishedDate": "2017-09-25",
          "description": "The late 16th century and the first half of the 17th century saw a final resurgence of the concept of Fortuna. Shortly thereafter, this goddess of chance and luck, who had survived for millennia, rapidly lost her cultural and intellectual relevance. This volume explores the late heyday and subsequent erasure of Fortuna. It examines vernacular traditions and confessional differences, analyses how the iconography and semantics of Fortuna motifs transformed, and traces the rise of complementary concepts such as those of probability, risk, fate and contingency. Thus, a multidisciplinary team of contributors sheds light on the surprising ways in which the end of Fortuna intersected with the rise of modernity.",
          "industryIdentifiers": [
            {
              "type": "ISBN_13",
              "identifier": "9783110455045"
            },
            {
              "type": "ISBN_10",
              "identifier": "3110455048"
            }
          ],
          "readingModes": {
            "text": false,
            "image": true
          },
          "pageCount": 224,
          "printType": "BOOK",
          "categories": [
            "History"
          ],
          "maturityRating": "NOT_MATURE",
          "allowAnonLogging": false,
          "contentVersion": "0.1.0.0.preview.1",
          "panelizationSummary": {
            "containsEpubBubbles": false,
            "containsImageBubbles": false
          },
          "imageLinks": {
            "smallThumbnail": "http://books.google.com/books/content?id=DSY-DwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
            "thumbnail": "http://books.google.com/books/content?id=DSY-DwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
          },
          "language": "en",
          "previewLink": "http://books.google.be/books?id=DSY-DwAAQBAJ&pg=PA80&dq=lipsius+admiranda+1598&hl=&cd=10&source=gbs_api",
          "infoLink": "https://play.google.com/store/books/details?id=DSY-DwAAQBAJ&source=gbs_api",
          "canonicalVolumeLink": "https://play.google.com/store/books/details?id=DSY-DwAAQBAJ"
        },
        "saleInfo": {
          "country": "BE",
          "saleability": "FOR_SALE",
          "isEbook": true,
          "listPrice": {
            "amount": 69.91,
            "currencyCode": "EUR"
          },
          "retailPrice": {
            "amount": 69.91,
            "currencyCode": "EUR"
          },
          "buyLink": "https://play.google.com/store/books/details?id=DSY-DwAAQBAJ&rdid=book-DSY-DwAAQBAJ&rdot=1&source=gbs_api",
          "offers": [
            {
              "finskyOfferType": 1,
              "listPrice": {
                "amountInMicros": 69910000,
                "currencyCode": "EUR"
              },
              "retailPrice": {
                "amountInMicros": 69910000,
                "currencyCode": "EUR"
              }
            }
          ]
        },
        "accessInfo": {
          "country": "BE",
          "viewability": "PARTIAL",
          "embeddable": true,
          "publicDomain": false,
          "textToSpeechPermission": "ALLOWED",
          "epub": {
            "isAvailable": false
          },
          "pdf": {
            "isAvailable": true,
            "acsTokenLink": "http://books.google.be/books/download/The_End_of_Fortuna_and_the_Rise_of_Moder-sample-pdf.acsm?id=DSY-DwAAQBAJ&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
          },
          "webReaderLink": "http://play.google.com/books/reader?id=DSY-DwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
          "accessViewStatus": "SAMPLE",
          "quoteSharingAllowed": false
        },
        "searchInfo": {
          "textSnippet": "... together with a discussion of tactics and defence systems.70 In his book on the cultural history of the Roman Empire, \u003cb\u003eAdmiranda\u003c/b\u003e sive De Magnitudine Romana (Antwerp \u003cb\u003e1598\u003c/b\u003e), \u003cb\u003eLipsius\u003c/b\u003e praised the military and political virtues of the&nbsp;..."
        }
      }
    ]
  }
`)

func TestGoogleBooksParse0(t *testing.T) {
	var response GoogleBooks
	err := json.Unmarshal(googlebooksJSON, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestGoogleBooksParse1(t *testing.T) {
	var response GoogleBooks
	_ = json.Unmarshal(googlebooksJSON, &response)
	result := response.TotalItems
	expected := 872
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%v]\nExpected: \n[%v]\n", result, expected))
	}
}
