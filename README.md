# Analyseur de Vulnérabilités WordPress

Ce projet a pour objectif de créer une plateforme web capable d’analyser des sites WordPress afin de détecter des failles de sécurité à l’aide d’un script exécuté dans un environnement Kali Linux et basé sur l’outil WPScan. Une fois l’analyse terminée, un rapport synthétique est généré, pouvant être exporté en PDF et envoyé par email.

## Présentation

Ce projet offre une solution centralisée pour réaliser des scans de sécurité sur des sites WordPress. Il exploite WPScan pour récupérer diverses informations techniques (version de WordPress, plugins, thèmes, etc.) et fournit aux utilisateurs un rapport détaillé.

## Objectifs

- Fournir une interface web intuitive pour lancer des analyses de vulnérabilités.
- Utiliser un script basé sur Kali Linux pour exécuter WPScan et vérifier les URL de sites WordPress.
- Générer un rapport synthétique de l’analyse incluant les informations collectées par WPScan.
- Permettre l’exportation du rapport au format PDF et son envoi par email.

## Fonctionnalités

- **Analyse de sites WordPress :** Exécution d’un scan via WPScan sur les URL ciblées.
- **Collecte d’informations :** Détection de la version de WordPress, des plugins, des thèmes et autres informations utiles.
- **Génération de rapports :** Création d’un rapport d’analyse synthétique à partir des résultats du scan.
- **Export PDF :** Possibilité d’exporter le rapport en PDF pour une diffusion facile.
- **Envoi par email :** Fonction intégrée pour envoyer directement le rapport par email.

## Prérequis

- **Kali Linux :** Le script est basé sur cet environnement, assurez-vous de l’avoir configuré.
- **WPScan :** L’outil doit être installé. Sous Kali, il peut être installé via la commande :
  
  ```bash
  sudo apt install wpscan
  ```

- **Ruby :** WPScan étant développé en Ruby, assurez-vous que Ruby et ses dépendances sont installés.


## Installation

1. **Cloner le dépôt :**  
   Clonez ce projet sur votre machine locale en utilisant `git clone`.

2. **Configurer l’environnement :**  
   - Installez WPScan et ses dépendances.
   - Configurez votre environnement Kali Linux.
   - Ajoutez votre token API WPScan dans les fichiers de configuration si nécessaire.

3. **Installation des dépendances supplémentaires :**  
   Selon votre stack (par exemple, pour l’interface web), installez les librairies ou frameworks requis (PHP, Python, Node.js, etc.).

## Utilisation

Pour lancer l’analyse d’une URL, vous pouvez utiliser l’interface web du projet ou directement exécuter le script de scan. Par exemple, en ligne de commande :

```bash
./run_scan.sh --url "mettre l'url du site"
```

Le script exécutera WPScan avec les options définies (extraction de plugins, thèmes, backups, etc.) et générera un rapport contenant toute l’analyse effectuée. Les options de WPScan (comme l’énumération des plugins vulnérables avec `-e vp` ou des thèmes vulnérables avec `-e vt`) peuvent être personnalisées en fonction de vos besoins.

## Rapport et Exportation

Après le scan, le rapport généré comprend :

- Des informations sur la configuration du site (version, headers, etc.)
- La liste des plugins et thèmes installés et leurs éventuelles vulnérabilités
- Des recommandations pour sécuriser le site

Le rapport peut être :
- **Exporté au format PDF** pour faciliter le partage et l’archivage.
- **Envoyé par email** directement via l’interface web, une fois que les paramètres de messagerie ont été configurés.

## Équipe Projet

L’équipe se compose de neuf membres répartis selon les filières et niveaux :

- **B1 Info :** Emmanuel, Ilyes, Katia, Julien
- **B1 Cyber :** Maxime, Léna
- **B3 Cyber :** Hugo, Colin
- **B3 Info :** Frederick

Chaque membre apporte son expertise, que ce soit en informatique ou en cybersécurité, afin de garantir la qualité et la robustesse de l’application.

## Licence

Mettre la licence MIT
